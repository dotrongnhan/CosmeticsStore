<template>
  <div class="header-wrapicon2" @mouseover="toggleUserDropdown" @mouseleave="toggleUserUp">
      <router-link :to="isLoginSuccess ? '/user' : '/login'">
            <img
              :src="user?.avatar ? user.avatar : defaultAvatar"
              class="header-icon1"
              style="width: 27px; border-radius: 50%"
              alt="Avatar"
            />
      </router-link>     
    <div
      class="header-cart header-dropdown"
      :class="{ 'show-header-dropdown': isShowUserDropdown && isLoginSuccess }"
    >
    <div class="dropdown-content">
    <router-link to="/user/profile">Profile</router-link>
    <router-link to="/user/settings">Settings</router-link>
    <router-link to="/" @click="logout()">Sign out</router-link>
  </div>
    </div>
  </div>
</template>

<script>
import defaultAvatar from "@/assets/images/icons/icon-header-01.png";
import { mapState, mapMutations, mapActions } from "vuex";
import { currency } from "@/utils/currency";

export default {
  name: "HeaderUserDropdown",
  data() {
      return { 
          defaultAvatar
      }
  },
  computed: {
    ...mapState("users", ["isShowUserDropdown", "isLoginSuccess", "user"]),
  },

  methods: {
    ...mapActions("users",["getUserExits"]),
    ...mapMutations("users", ["updateUser"]),
    toggleUserDropdown() {
      this.$store.commit("users/setShowUserDropdown", true);
    },
    toggleUserUp() {
      this.$store.commit("users/setShowUserDropdown", false);
    },
    logout() {
      this.$store.dispatch("users/logout");
    },
    currency,
  },
};
</script>

<style scoped>
.header-cart {
    margin-right: -30px;
    width: 140px;
}

.dropdown-content a {
  color: black;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
}

.dropdown-content a:hover {
        background-color: lightgrey;
}
</style>
