<template>
  <div>
   <h1>Welcome</h1>
    <router-link to="/login" v-if="!isLoggedIn"><button>Login</button></router-link>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  computed: {
    ...mapGetters(['isLoggedIn'])
  }
}
</script>