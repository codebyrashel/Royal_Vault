<template>
  <div>
    <h2>Register</h2>
    <input v-model="email" placeholder="Email" />
    <input v-model="password" type="password" placeholder="Password" />
    <button @click="register">Register</button>
  </div>
</template>

<script>
import api from "../api/api";

export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    async register() {
      await api.post("/auth/register", {
        email: this.email,
        password: this.password,
      });

      this.$router.push("/login");
    },
  },
};
</script>
