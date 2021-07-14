<template>
  <MainHeader />
  <router-view v-slot="{ Component }">
    <transition name="fade">
      <component :is="Component" />
    </transition>
  </router-view>
  <MainFooter />
</template>

<script>
import MainHeader from "@/layout/MainHeader.vue";
import MainFooter from "@/layout/MainFooter.vue";

export default {
  name: "App",
  components: {
    MainHeader,
    MainFooter,
  },
};
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
