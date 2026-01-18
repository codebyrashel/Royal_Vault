<template>
  <div class="auth-page">
    <h1>Log in to Royal Vault</h1>

    <form @submit.prevent="handleLogin" class="auth-form">
      <div class="field">
        <label>Email</label>
        <input v-model="email" type="email" required />
      </div>

      <div class="field">
        <label>Login password</label>
        <input v-model="loginPassword" type="password" required />
      </div>

      <div class="field">
        <label>Master password</label>
        <input v-model="masterPassword" type="password" required />
        <small>
          Used only to decrypt your vault locally. Never sent to the server.
        </small>
      </div>

      <button type="submit" :disabled="loading">
        {{ loading ? 'Logging in...' : 'Log in' }}
      </button>

      <p v-if="error" class="error">{{ error }}</p>

      <p class="switch">
        Need an account?
        <a @click.prevent="$router.push('/signup')">Sign up</a>
      </p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { login } from '../api/auth';
import { useAuthStore } from '../stores/authStore';
import {
  deriveKeyFromPassword,
  importVaultKey,
  decryptText,
} from '../utils/crypto';

const router = useRouter();
const { setAuth, setVaultKey } = useAuthStore();

const email = ref('');
const loginPassword = ref('');
const masterPassword = ref('');
const loading = ref(false);
const error = ref('');

async function handleLogin() {
  error.value = '';
  loading.value = true;

  try {
    if (!email.value) {
      throw new Error('Email is required');
    }
    if (!loginPassword.value) {
      throw new Error('Login password is required');
    }
    if (!masterPassword.value) {
      throw new Error('Master password is required');
    }

    // 1. Call backend login
    const res = await login({
      email: email.value,
      password: loginPassword.value,
    });

    try {
      // 2. Decode salt and derive master key again
      const saltBytes = Uint8Array.from(
        window.atob(res.salt),
        (c) => c.charCodeAt(0)
      );
      const { key: masterKey } = await deriveKeyFromPassword(
        masterPassword.value,
        saltBytes
      );

      // 3. Parse encryptedVaultKey payload and decrypt vault key bytes
      const encryptedVaultKeyPayload = JSON.parse(res.encryptedVaultKey);
      const vaultKeyBytesBase64 = await decryptText(
        masterKey,
        encryptedVaultKeyPayload
      );
      const vaultKeyBytes = Uint8Array.from(
        window.atob(vaultKeyBytesBase64),
        (c) => c.charCodeAt(0)
      );
      const vaultKey = await importVaultKey(vaultKeyBytes);

      // 4. Save auth token and vault key in store
      setAuth(res.token, res.email);
      setVaultKey(vaultKey);

      // 5. Redirect to dashboard
      router.push('/dashboard');
    } catch (decryptError: any) {
      // This usually means master password was wrong or vault data is corrupted
      throw new Error(
        'Master password is incorrect or vault data could not be decrypted'
      );
    }
  } catch (e: any) {
    error.value = e?.message ?? 'Login failed';
  } finally {
    loading.value = false;
  }
}
</script>

<style scoped>
.auth-page {
  max-width: 480px;
  margin: 2rem auto;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
.field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}
input {
  padding: 0.5rem;
  border-radius: 0.25rem;
  border: 1px solid #d1d5db;
}
small {
  font-size: 0.8rem;
  color: #6b7280;
}
button {
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
  border: none;
  cursor: pointer;
  background-color: #2563eb;
  color: white;
}
button:disabled {
  opacity: 0.6;
  cursor: default;
}
.error {
  color: #b91c1c;
}
.switch {
  font-size: 0.9rem;
}
.switch a {
  color: #2563eb;
  cursor: pointer;
}
</style>