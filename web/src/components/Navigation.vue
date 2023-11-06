<template>
  <div class="fixed bottom-0 top-0 left-0 flex items-center justify-items-center">
    <div class="surface-card backdrop-blur p-1  rounded-lg ml-1 drop-shadow-lg">
      <div v-if="state.tagMode" @click="back" v-tooltip="'返回'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i class="pi pi-arrow-left"></i>
      </div>
      <div @click="edit" v-tooltip="'编辑'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i class="pi pi-file-edit"></i>
      </div>
      <div @click="addClick" v-tooltip="'添加'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i  class="pi pi-plus"></i>
      </div>
      <div v-tooltip="'删除'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i class="pi pi-trash"></i>
      </div>
      <div v-if="!state.tagMode" v-tooltip="'导出excel'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i class="pi pi-file-export"></i>
      </div>
      <div v-if="!state.tagMode" v-tooltip="'导入excel'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i class="pi pi-file-import"></i>
      </div>
      <div v-if="!state.tagMode" @click="tags" v-tooltip="'标签管理'" class="cursor-pointer p-3 hover:bg-orange-300">
        <i class="pi pi-tags"></i>
      </div>

    </div>
    <Dialog  v-model:visible="visible" modal header="添加流行语" :style="{ width: '25vw' }" >
      <InputText placeholder="请输入流行语名称" class="w-full my-2 p-2 border-1"/>
      <MultiSelect placeholder="流行语名称标签" class="w-full my-2 border-1"/>
      <Textarea placeholder="请输入流行语释义" class="w-full my-2 p-2 border-1"/>
      <template #footer>
        <Button class="border-1 p-2 hover:bg-blue-300" label="添加" @click="visible = false" autofocus />
      </template>
    </Dialog>
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {useNavigationStore} from "../state.ts";
import {useRouter} from "vue-router";
const state=useNavigationStore()
const router=useRouter()

const visible=ref(false)
const addClick=function (){
  visible.value=true;
}
const edit=function (){
  state.editMode=!state.editMode;
}
const tags=function (){
  state.tagMode=true;
  router.push("/tag")
}


const back=function () {
  state.tagMode=false
  router.push("/")
}
</script>