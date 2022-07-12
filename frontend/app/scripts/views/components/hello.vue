<template>
  <div class="hello">
    <ui-button icon="add" @click="$alert('Hello BalmUI')">Click Me</ui-button>
  </div>
</template>

<script>
export default {
  name: 'hello',
  setup() {
    return {};
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
