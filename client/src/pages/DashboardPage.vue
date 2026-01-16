<template>
  <div class="dashboard-shell">
    <header class="dashboard-header">
      <div class="logo">Royal Vault</div>
      <nav class="nav-links">
        <span class="nav-item">Dashboard</span>
        <span class="nav-item muted">Credentials (coming)</span>
      </nav>
    </header>

    <main class="dashboard-main">
      <h1>Dashboard</h1>
      <p>
        This is a placeholder for your secure credential dashboard.
        In future steps, this will show your vault, folders, and credentials.
      </p>
    </main>
  </div>
</template>

<script setup lang="ts">
// Placeholder dashboard, logic will be added later
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

.dashboard-main {
  padding: 2rem;
  max-width: 960px;
  margin: 0 auto;
}

.dashboard-main h1 {
  font-size: 1.75rem;
  margin-bottom: 0.75rem;
  color: #111827;
}

.dashboard-main p {
  font-size: 0.98rem;
  color: #4b5563;
}
</style>