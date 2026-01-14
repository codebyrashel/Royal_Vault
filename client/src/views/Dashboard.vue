<template>
  <div>
    <h2>Dashboard</h2>
    <button @click="loadProfile">Load Profile</button>
    <pre>{{ user }}</pre>
  </div>
</template>

<script>
import api from "../api/api";

export default {
  data() {
    return {
      user: null,
    };
  },
  methods: {
    async loadProfile() {
      const res = await api.get("/vault/me");
      this.user = res.data;
    },
  },
};
</script>
