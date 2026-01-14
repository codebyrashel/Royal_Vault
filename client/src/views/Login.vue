<template>
  <div>
    <h2>Login</h2>
    <input v-model="email" placeholder="Email" />
    <input v-model="password" type="password" placeholder="Password" />
    <button @click="login">Login</button>
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
    async login() {
      const res = await api.post("/auth/login", {
        email: this.email,
        password: this.password,
      });

      localStorage.setItem("token", res.data.token);
      this.$router.push("/dashboard");
    },
  },
};
</script>
