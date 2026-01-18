// src/stores/authStore.ts
import { ref } from 'vue';

const token = ref<string | null>(null);
const vaultKey = ref<CryptoKey | null>(null);
const userEmail = ref<string | null>(null);

export function useAuthStore() {
  function setAuth(newToken: string, email: string) {
    token.value = newToken;
    userEmail.value = email;
  }

  function clearAuth() {
    token.value = null;
    userEmail.value = null;
    vaultKey.value = null;
  }

  function setVaultKey(key: CryptoKey) {
    vaultKey.value = key;
  }

   function isAuthenticated(): boolean {
    return !!token.value && !!vaultKey.value;
  }

  return {
    token,
    userEmail,
    vaultKey,
    setAuth,
    clearAuth,
    setVaultKey,
    isAuthenticated,
  };
}