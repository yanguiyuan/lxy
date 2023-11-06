<script setup lang="ts">
import {useNavigationStore} from "../state.ts";

const props=defineProps({
  word:String,
  tags:Array<String>,
  comment:String,
})
const state=useNavigationStore()
</script>

<template>
  <li class="bg-white shadow rounded p-4">
    <div class="flex items-center justify-items-start mb-2">
      <input v-if="state.editMode" type="checkbox" class="form-checkbox text-indigo-600 rounded">
      <h2 class="text-lg font-bold ml-2">{{ props.word }}</h2>
    </div>
    <Tag class="mr-1" v-for="tag in props.tags" rounded :style="{ background: 'linear-gradient(-225deg,#AC32E4 0%,#7918F2 48%,#4801FF 100%)' }">
      <div class="flex align-items-center gap-2">
        <span class="text-base text-xs">{{ tag }}</span>
        <i v-if="state.editMode" class="pi pi-times text-xs cursor-pointer"></i>
      </div>
    </Tag>
    <Tag v-if="state.editMode" class="cursor-pointer" :style="{ background: 'linear-gradient(-225deg,#AC32E4 0%,#7918F2 48%,#4801FF 100%)' }">
      <i class="pi pi-plus"></i>
    </Tag>
    <p class="text-gray-700">{{props.comment}}</p>
  </li>
</template>

<style scoped>

</style>