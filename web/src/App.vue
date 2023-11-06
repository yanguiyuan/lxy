<template>
  <div style="margin: 0 200px 0 200px">
    <RouterView/>
  <Navigation/>
  </div>
</template>
<script setup lang="ts">
import Navigation from "./components/Navigation.vue";
</script>