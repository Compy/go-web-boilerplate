<template>
  <v-container fluid fill-height id="login-page">
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-1">
          <v-form @submit.prevent="login">
            <v-card-text>
              <v-alert type="error" v-if="error">{{ error }}</v-alert>
              <v-text-field
                outlined
                dense
                autofocus
                name="login"
                label="Email"
                type="text"
                v-model="email"
              ></v-text-field>
              <v-text-field
                outlined
                dense
                id="password"
                name="password"
                label="Password"
                type="password"
                v-model="password"
              ></v-text-field>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" type="submit">Login</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>
<script>
import { mapMutations } from "vuex";
export default {
  data: () => {
    return {
      email: "",
      password: "",
      error: false,
    };
  },
  methods: {
    ...mapMutations(["setUser", "setToken"]),
    async login(e) {
      e.preventDefault();
      const response = await fetch("http://localhost:8080/api/token", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: this.email,
          password: this.password,
        }),
      });
      const { token, error } = await response.json();
      if (token === undefined || !response.ok) {
        this.error = error;
        return;
      }
      this.setToken(token);
      this.$router.push("/");
    },
  },
};
</script>
<style scoped>
#login-page {
  background-color: #fafafa;
}
</style>