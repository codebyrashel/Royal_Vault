<template>
  <div class="auth-page">
    <h1>Create your Royal Vault account</h1>

    <form @submit.prevent="handleSignup" class="auth-form">
      <div class="field">
        <label>Email</label>
        <input v-model="email" type="email" required />
      </div>

      <div class="field">
        <label>Login password</label>
        <input v-model="loginPassword" type="password" required />
        <small>This is used only for logging into the service.</small>
      </div>

      <div class="field">
        <label>Master password</label>
        <input v-model="masterPassword" type="password" required />
        <small>
          This protects your vault. It is never sent to the server.
          If you lose it, your encrypted data cannot be recovered.
        </small>
      </div>

      <button type="submit" :disabled="loading">
        {{ loading ? 'Creating account...' : 'Sign up' }}
      </button>

      <p v-if="error" class="error">{{ error }}</p>
      <p v-if="success" class="success">{{ success }}</p>

      <p class="switch">
        Already have an account?
        <a @click.prevent="$router.push('/login')">Log in</a>
      </p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { signup } from '../api/auth';
import {
  deriveKeyFromPassword,
  generateVaultKey,
  exportKey,
  encryptText,
} from '../utils/crypto';

const router = useRouter();

const email = ref('');
const loginPassword = ref('');
const masterPassword = ref('');
const loading = ref(false);
const error = ref('');
const success = ref('');

async function handleSignup() {
  error.value = '';
  success.value = '';
  loading.value = true;

  try {
    if (!email.value) {
      throw new Error('Email is required');
    }
    if (!loginPassword.value || loginPassword.value.length < 8) {
      throw new Error('Login password must be at least 8 characters long');
    }
    if (!masterPassword.value) {
      throw new Error('Master password is required');
    }

    // 1. Derive master key and get salt
    const { key: masterKey, salt } = await deriveKeyFromPassword(
      masterPassword.value
    );

    // 2. Generate vault key and export raw bytes
    const vaultKey = await generateVaultKey();
    const vaultKeyBytes = await exportKey(vaultKey);

    // 3. Encrypt vault key bytes with master key
    const vaultKeyBytesBase64 = window.btoa(
      String.fromCharCode(...Array.from(vaultKeyBytes))
    );
    const encryptedVaultKeyPayload = await encryptText(
      masterKey,
      vaultKeyBytesBase64
    );
    const encryptedVaultKeyString = JSON.stringify(encryptedVaultKeyPayload);

    // 4. Prepare payload for backend
    const saltBase64 = window.btoa(
      String.fromCharCode(...Array.from(salt))
    );

    await signup({
      email: email.value,
      password: loginPassword.value,
      encryptedVaultKey: encryptedVaultKeyString,
      salt: saltBase64,
    });

    success.value = 'Account created. You can now log in.';
    setTimeout(() => {
      router.push('/login');
    }, 1000);
  } catch (e: any) {
    error.value = e?.message ?? 'Signup failed';
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
.success {
  color: #15803d;
}
.switch {
  font-size: 0.9rem;
}
.switch a {
  color: #2563eb;
  cursor: pointer;
}
</style>