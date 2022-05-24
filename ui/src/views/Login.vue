<template>
    <div>
        <h1>LOGIN</h1>
        <div v-if="error">
        {{ error }}
        </div>
        <form @submit.prevent="login">
            <div>
            <input v-model="email" placeholder="email">
            </div>
            <div>
                <input v-model="password" placeholder="password" type="password">
            </div>
            <button type="submit">Login</button>
        </form>
    </div>
</template>
<script>
import { mapMutations } from 'vuex'
export default {
    data: () => {
        return {
            email: "",
            password: "",
            error: false
        };
    },
    methods: {
        ...mapMutations(["setUser", "setToken"]),
        async login(e) {
            e.preventDefault()
            const response = await fetch("http://localhost:8080/api/token", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    email: this.email,
                    password: this.password,
                }),
            })
            const { token, error } = await response.json()
            if (token === undefined || !response.ok) {
                this.error = error
                return
            }
            this.setToken(token)
            this.$router.push("/")
        },
    },
};
</script>
