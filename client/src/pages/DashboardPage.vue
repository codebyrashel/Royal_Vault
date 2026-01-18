<template>
  <div class="dashboard-shell">
    <header class="dashboard-header">
      <div class="logo">Royal Vault</div>
      <nav class="nav-links">
        <span class="nav-item">Dashboard</span>
        <span class="nav-item muted">Credentials</span>
      </nav>
      <div class="user-info" v-if="auth.userEmail">
        <span>{{ auth.userEmail }}</span>
        <button class="small-btn" @click="handleLogout">Log out</button>
      </div>
    </header>

    <main class="dashboard-main">
      <section v-if="!auth.isAuthenticated()" class="warning">
        <p>
          You are not fully authenticated (missing token or vault key).
          Please <router-link to="/login">log in</router-link> again.
        </p>
      </section>

      <section v-else>
        <h1>Your Credentials</h1>

        <div class="actions">
          <button @click="showCreateForm = !showCreateForm">
            {{ showCreateForm ? 'Cancel' : 'Add Credential' }}
          </button>
        </div>

        <form v-if="showCreateForm" @submit.prevent="handleCreate" class="create-form">
          <div class="field">
            <label>Title</label>
            <input v-model="form.title" required />
          </div>
          <div class="field">
            <label>URL</label>
            <input v-model="form.url" placeholder="https://example.com" />
          </div>
          <div class="field">
            <label>Username</label>
            <input v-model="form.username" required />
          </div>
          <div class="field">
            <label>Password</label>
            <input v-model="form.password" required />
          </div>
          <div class="field">
            <label>Notes</label>
            <textarea v-model="form.notes" rows="2" />
          </div>
          <button type="submit" :disabled="loading">
            {{ loading ? 'Saving...' : 'Save credential' }}
          </button>
          <p v-if="formError" class="error">{{ formError }}</p>
        </form>

        <section class="list-section">
          <h2>Stored credentials</h2>
          <p v-if="loading && !credentials.length">Loading...</p>
          <p v-if="!loading && !credentials.length">No credentials yet.</p>
          <p v-if="loadError" class="error">{{ loadError }}</p>

          <table v-if="credentials.length" class="cred-table">
            <thead>
              <tr>
                <th>Title</th>
                <th>URL</th>
                <th>Username</th>
                <th>Password</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="cred in credentials" :key="cred.id">
                <td>{{ cred.title }}</td>
                <td>
                  <a v-if="cred.url" :href="cred.url" target="_blank" rel="noreferrer">
                    {{ cred.url }}
                  </a>
                </td>
                <td>{{ cred.decryptedUsername }}</td>
                <td>
                  <span v-if="cred.showPassword">{{ cred.decryptedPassword }}</span>
                  <span v-else>••••••••</span>
                  <button class="small-btn" @click="togglePassword(cred)">
                    {{ cred.showPassword ? 'Hide' : 'Show' }}
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </section>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/authStore';
import { fetchCredentials, createCredential, type CredentialResponse } from '../api/credentials';
import { decryptText, encryptText } from '../utils/crypto';

const router = useRouter();
const auth = useAuthStore();

interface DecryptedCredential extends CredentialResponse {
  decryptedUsername: string;
  decryptedPassword: string;
  decryptedNotes?: string | null;
  showPassword: boolean;
}

const credentials = ref<DecryptedCredential[]>([]);
const loading = ref(false);
const loadError = ref('');
const formError = ref('');
const showCreateForm = ref(false);

const form = reactive({
  title: '',
  url: '',
  username: '',
  password: '',
  notes: '',
});

function resetForm() {
  form.title = '';
  form.url = '';
  form.username = '';
  form.password = '';
  form.notes = '';
}

function handleLogout() {
  auth.clearAuth();
  router.push('/login');
}

function togglePassword(cred: DecryptedCredential) {
  cred.showPassword = !cred.showPassword;
}

async function loadCredentials() {
  if (!auth.token.value || !auth.vaultKey.value) {
    return;
  }

  loading.value = true;
  loadError.value = '';

  try {
    const res = await fetchCredentials(auth.token.value);
    const vaultKey = auth.vaultKey.value;

    const decryptedList: DecryptedCredential[] = [];

    for (const c of res) {
      const username = await decryptText(vaultKey, JSON.parse(c.encryptedUsername));
      const password = await decryptText(vaultKey, JSON.parse(c.encryptedPassword));
      let notes: string | null = null;
      if (c.encryptedNotes) {
        notes = await decryptText(vaultKey, JSON.parse(c.encryptedNotes));
      }

      decryptedList.push({
        ...c,
        decryptedUsername: username,
        decryptedPassword: password,
        decryptedNotes: notes,
        showPassword: false,
      });
    }

    credentials.value = decryptedList;
  } catch (e: any) {
    loadError.value = e?.message ?? 'Failed to load credentials';
  } finally {
    loading.value = false;
  }
}

async function handleCreate() {
  formError.value = '';
  if (!auth.token.value || !auth.vaultKey.value) {
    formError.value = 'You are not authenticated. Please log in again.';
    return;
  }
  if (!form.title || !form.username || !form.password) {
    formError.value = 'Title, username, and password are required.';
    return;
  }

  loading.value = true;

  try {
    const vaultKey = auth.vaultKey.value;
    // Encrypt sensitive fields with vault key
    const encryptedUsername = await encryptText(vaultKey, form.username);
    const encryptedPassword = await encryptText(vaultKey, form.password);
    const encryptedNotes = form.notes
      ? await encryptText(vaultKey, form.notes)
      : null;

    const payload = {
      title: form.title,
      url: form.url || null,
      encryptedUsername: JSON.stringify(encryptedUsername),
      encryptedPassword: JSON.stringify(encryptedPassword),
      encryptedNotes: encryptedNotes ? JSON.stringify(encryptedNotes) : null,
    };

    const created = await createCredential(auth.token.value, payload);

    // Decrypt just-created credential and append
    const decUsername = await decryptText(vaultKey, JSON.parse(created.encryptedUsername));
    const decPassword = await decryptText(vaultKey, JSON.parse(created.encryptedPassword));
    let decNotes: string | null = null;
    if (created.encryptedNotes) {
      decNotes = await decryptText(vaultKey, JSON.parse(created.encryptedNotes));
    }

    credentials.value.push({
      ...created,
      decryptedUsername: decUsername,
      decryptedPassword: decPassword,
      decryptedNotes: decNotes,
      showPassword: false,
    });

    resetForm();
    showCreateForm.value = false;
  } catch (e: any) {
    formError.value = e?.message ?? 'Failed to create credential';
  } finally {
    loading.value = false;
  }
}

onMounted(() => {
  if (auth.isAuthenticated()) {
    loadCredentials();
  }
});
</script>

<style scoped>
.dashboard-shell {
  min-height: 100vh;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}
.dashboard-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 2rem;
  border-bottom: 1px solid #e5e7eb;
}
.logo {
  font-weight: 700;
  font-size: 1.25rem;
}
.nav-links {
  display: flex;
  gap: 1rem;
}
.nav-item {
  font-size: 0.95rem;
  color: #4b5563;
}
.nav-item.muted {
  color: #9ca3af;
}
.user-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 0.9rem;
}
.small-btn {
  padding: 0.25rem 0.5rem;
  font-size: 0.8rem;
  border-radius: 0.25rem;
  border: 1px solid #d1d5db;
  background: white;
  cursor: pointer;
}
.dashboard-main {
  padding: 2rem;
  max-width: 960px;
  margin: 0 auto;
}
h1 {
  font-size: 1.75rem;
  margin-bottom: 0.75rem;
  color: #111827;
}
h2 {
  margin-top: 2rem;
  margin-bottom: 0.5rem;
  font-size: 1.25rem;
}
.actions {
  margin: 1rem 0;
}
.create-form {
  margin-bottom: 2rem;
  padding: 1rem;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  background: #f9fafb;
}
.field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  margin-bottom: 0.75rem;
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
  margin-top: 0.5rem;
  color: #b91c1c;
}
.warning {
  padding: 1rem;
  border-radius: 0.5rem;
  background-color: #fef3c7;
  border: 1px solid #facc15;
}
.cred-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 0.5rem;
}
.cred-table th,
.cred-table td {
  border: 1px solid #e5e7eb;
  padding: 0.5rem;
  font-size: 0.9rem;
}
.cred-table th {
  background-color: #f3f4f6;
  text-align: left;
}
</style>