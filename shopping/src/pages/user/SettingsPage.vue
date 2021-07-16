<template>
  <div class="mb-3">
    <ul class="nav nav-tabs" id="myTab" role="tablist">
      <li @click="componentSelected = 'ManageProducts'" class="nav-item">
        <button :class="[componentSelected === 'ManageProducts' ? 'active nav-link' : 'nav-link']">Products</button>
      </li>
      <li @click="componentSelected = 'ManageUser'" class="nav-item m-text6">
        <button :class="[componentSelected === 'ManageUser' ? 'nav-link active' : 'nav-link']">User</button>
      </li>
    </ul>
  </div>
  <component :is="componentSelected"/>
</template>

<script>
import ManageProducts from "../../components/ManageProducts";

export default {
  name: "SettingsPage",
  components: {
    ManageProducts
  },
  data() {
    return {
      componentSelected: "ManageProducts"
    }
  }
};
</script>

<style scoped>
</style>