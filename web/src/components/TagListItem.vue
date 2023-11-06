<script setup lang="ts">
import {useNavigationStore} from "../state.ts";

const state=useNavigationStore()
const props=defineProps({
  tag:String,
  tagType:String,
})
</script>

<template>
  <li class="bg-white shadow rounded p-4 mb-1">
    <div class="flex items-center justify-items-start mb-2">
      <input v-if="state.editMode" type="checkbox" class="form-checkbox text-indigo-600 rounded">
      <h2 class="text-lg font-bold ml-2">{{ props.tag }}</h2>
      <p class="text-gray-700 ml-5">标签类型：{{props.tagType}}</p>
    </div>

  </li>
</template>

<style scoped>

</style>