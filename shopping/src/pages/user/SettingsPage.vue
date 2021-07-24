<template>
  <div class="mb-3">
    <ul class="nav nav-tabs" id="myTab" role="tablist">
      <li @click="componentSelected = 'ManageProducts'" class="nav-item">
        <button :class="[componentSelected === 'ManageProducts' ? 'active nav-link' : 'nav-link']">Products</button>
      </li>
      <li @click="componentSelected = 'ManageUsers'" class="nav-item m-text6">
        <button :class="[componentSelected === 'ManageUsers' ? 'nav-link active' : 'nav-link']">User</button>
      </li>
      <li @click="componentSelected = 'ManageOrders'" class="nav-item m-text6">
        <button :class="[componentSelected === 'ManageOrders' ? 'nav-link active' : 'nav-link']">Order</button>
      </li>
    </ul>
  </div>
  <component :is="componentSelected"/>
</template>

<script>
import ManageProducts from "../../components/ManageProducts";
import ManageUsers from "../../components/ManageUsers";
import ManageOrders from "../../components/ManageOrders";
import {mapState} from "vuex";

export default {
  name: "SettingsPage",
  components: {
    ManageProducts,
    ManageUsers,
    ManageOrders
  },
  data() {
    return {
      componentSelected: "ManageProducts"
    }
  },
  created() {
    if(this.user.role_id !== 1) {
      this.$router.push("/")
    }
  },
  computed: {
    ...mapState("users", ["user"])
  }
};
</script>

<style scoped>
</style>