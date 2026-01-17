<template>
  <div class="crypto-test">
    <h1>Crypto Test</h1>
    <div class="field">
      <label>Master Password</label>
      <input v-model="masterPassword" type="password" />
    </div>
    <div class="field">
      <label>Plaintext</label>
      <textarea v-model="plaintext" rows="3" />
    </div>
    <button @click="runTest" :disabled="isRunning">
      Run Encryption Test
    </button>

    <div v-if="error" class="error">
      <strong>Error:</strong> {{ error }}
    </div>

    <div v-if="result">
      <h2>Result</h2>
      <pre>{{ result }}</pre>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import {
  deriveKeyFromPassword,
  generateVaultKey,
  exportKey,
  importVaultKey,
  encryptText,
  decryptText,
} from '../utils/crypto';

const masterPassword = ref('');
const plaintext = ref('example secret');
const isRunning = ref(false);
const result = ref('');
const error = ref('');

async function runTest() {
  error.value = '';
  result.value = '';
  isRunning.value = true;

  try {
    if (!masterPassword.value) {
      throw new Error('Master password is required');
    }

    // 1. Derive master key
    const { key: masterKey, salt } = await deriveKeyFromPassword(
      masterPassword.value
    );

    // 2. Generate vault key and export it
    const vaultKey = await generateVaultKey();
    const vaultKeyBytes = await exportKey(vaultKey);

    // 3. Encrypt the vault key with the master key
    const vaultKeyEncrypted = await encryptText(masterKey, btoa(String.fromCharCode(...vaultKeyBytes)));

    // 4. Decrypt the vault key using the same master password and salt
    const { key: masterKey2 } = await deriveKeyFromPassword(
      masterPassword.value,
      salt
    );
    const decryptedVaultKeyBytesBase64 = await decryptText(
      masterKey2,
      vaultKeyEncrypted
    );
    const decryptedVaultKeyBytes = Uint8Array.from(
      atob(decryptedVaultKeyBytesBase64),
      (c) => c.charCodeAt(0)
    );
    const importedVaultKey = await importVaultKey(decryptedVaultKeyBytes);

    // 5. Use the imported vault key to encrypt/decrypt the plaintext
    const encryptedPayload = await encryptText(importedVaultKey, plaintext.value);
    const decryptedText = await decryptText(importedVaultKey, encryptedPayload);

    result.value = JSON.stringify(
      {
        salt: btoa(String.fromCharCode(...salt)),
        vaultKeyEncrypted,
        encryptedPayload,
        decryptedText,
      },
      null,
      2
    );
  } catch (e: any) {
    error.value = e?.message ?? String(e);
  } finally {
    isRunning.value = false;
  }
}
</script>

<style scoped>
.crypto-test {
  max-width: 600px;
  margin: 2rem auto;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}
.field {
  margin-bottom: 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}
input,
textarea {
  padding: 0.5rem;
  border-radius: 0.25rem;
  border: 1px solid #d1d5db;
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
  margin-top: 1rem;
  color: #b91c1c;
}
pre {
  margin-top: 1rem;
  background: #f3f4f6;
  padding: 0.75rem;
  border-radius: 0.25rem;
  font-size: 0.85rem;
  overflow-x: auto;
}
</style>