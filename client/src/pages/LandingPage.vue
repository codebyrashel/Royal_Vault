<template>
  <div class="app-shell">
    <header class="app-header">
      <div class="logo">Royal Vault</div>
      <nav class="nav-links">
        <a href="#how-it-works">How it works</a>
        <a href="#security">Security</a>
        <a href="#faq">FAQ</a>
      </nav>
      <button class="primary-btn" @click="$router.push('/dashboard')">
        Go to Dashboard
      </button>
    </header>

    <main class="app-main">
      <section class="hero">
        <div class="hero-content">
          <h1>One master key for all your credentials</h1>
          <p>
            Royal Vault is a zero-knowledge password manager that keeps your
            accounts organized and your secrets encrypted - even from us.
          </p>
          <div class="hero-actions">
            <button class="primary-btn" @click="$router.push('/dashboard')">
              Get started
            </button>
            <button class="secondary-btn">
              Learn more
            </button>
          </div>
        </div>
      </section>

      <section id="how-it-works" class="section">
        <h2>How it works</h2>
        <ol>
          <li>Create an account with your email.</li>
          <li>Set a master password that never leaves your device.</li>
          <li>Store and organize all your credentials in one encrypted vault.</li>
        </ol>
      </section>

      <section id="security" class="section">
        <h2>Security in plain English</h2>
        <p>
          Your passwords are encrypted in your browser before they are sent to
          our servers. We only store encrypted data, never the master key or
          decrypted secrets.
        </p>
      </section>

      <section id="faq" class="section">
        <h2>FAQ</h2>
        <p>More details will be added here as the product evolves.</p>
      </section>
    </main>

    <footer class="app-footer">
      <p>&copy; {{ new Date().getFullYear() }} Royal Vault. All rights reserved.</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
// No logic yet for the landing page
</script>

<style scoped>
.app-shell {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.app-header {
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

.nav-links a {
  text-decoration: none;
  color: #4b5563;
  font-size: 0.95rem;
}

.nav-links a:hover {
  color: #111827;
}

.primary-btn,
.secondary-btn {
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  border: none;
  cursor: pointer;
  font-size: 0.95rem;
}

.primary-btn {
  background-color: #2563eb;
  color: white;
}

.primary-btn:hover {
  background-color: #1d4ed8;
}

.secondary-btn {
  background-color: white;
  color: #374151;
  border: 1px solid #d1d5db;
}

.secondary-btn:hover {
  background-color: #f9fafb;
}

.app-main {
  flex: 1;
  padding: 2rem;
  max-width: 960px;
  margin: 0 auto;
}

.hero {
  padding: 2rem 0;
}

.hero-content{
  display: flex;
  flex-direction: column;
  align-items: center;
}


.hero-content h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: #111827;
}

.hero-content p {
  font-size: 1.05rem;
  color: #4b5563;
  max-width: 36rem;
}

.hero-actions {
  margin-top: 1.5rem;
  display: flex;
  gap: 1rem;
}

.section {
  margin-top: 3rem;
}

.section h2 {
  font-size: 1.5rem;
  margin-bottom: 0.75rem;
  color: #111827;
}

.section p,
.section ol {
  color: #4b5563;
  font-size: 0.98rem;
}

.app-footer {
  padding: 1rem 2rem;
  border-top: 1px solid #e5e7eb;
  text-align: center;
  font-size: 0.85rem;
  color: #6b7280;
}
</style>