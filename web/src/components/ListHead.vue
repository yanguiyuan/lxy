<script setup lang="ts">

import MultiSelect from "primevue/multiselect";
</script>

<template>
  <div class="mb-3">
         <span class="p-input-icon-left drop-shadow mb-1">
            <i class="pi pi-search" />
            <InputText class="pl-5" placeholder="Search" />
         </span>
    <MultiSelect
        filter optionLabel="name" placeholder="选择标签" :maxSelectedLabels="3"
        class="p-0 ml-2 w-4 drop-shadow"
        :pt="{
              root:{class:'p-0'},
              label:{class:'p-0 ml-1'},
              panel:{class:'mt-1 border-1'},
              filterContainer:{class:'border-1 rounded-md'},
              headerCheckboxContainer:{class:'border-1'}
            }"
    />
  </div>
</template>

<style scoped>

</style>