<template>
<div class="app-wrapper">
  <sidebar />
  <div class="main-container">
    <router-view></router-view>
  </div>
</div>
</template>

<script>
import sidebar from '@/views/components/sidebar';

export default {
  name: 'App',
  components: {
    sidebar
  },
  setup() {
    return {};
  }
};
</script>

<style scoped>
</style>