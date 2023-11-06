

<template>
  <div class="container mx-auto p-4 ">
  <h1 class="text-2xl font-bold mb-4">标签</h1>
    <div >
      <DataTable editMode="row"  @rowExpand="onRowExpand"  v-model:expandedRows="expandedRows" :value="tagData" dataKey="id" :pt="{

         wrapper:{class:'rounded-lg'},
         headerRow:{class:'border-1'},
         bodyRow:{class:'border-1'}
      }">
        <Column expander style="width: 5rem" />
        <Column field="tagType" header="标签类型"></Column>
        <Column :rowEditor="true" style="width: 10%; min-width: 8rem" bodyStyle="text-align:center"></Column>
        <template #expansion="slotProps">
          <div>
            <DataTable  v-model:selection="selectedProduct" :value="slotProps.data.tags">
              <Column  selectionMode="multiple" headerStyle="width: 3rem" :pt="{
                checkboxWrapper:{class:'border-1 border-gray-400 rounded'},
                headerCheckbox:{class:'border-1 border-gray-400'}
              }"></Column>
              <Column field="tag" header="标签"></Column>
            </DataTable>
          </div>
        </template>
      </DataTable>
    </div>
  </div>
</template>
<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {useNavigationStore} from "../state.ts";
const expandedRows = ref([]);
const selectedProduct = ref([]);
const tagData=ref([
  {
    id:1,
    tagType:"年份",
    tags:[
      {tag:"2007"},
      {tag:"2008"},
      {tag:"2009"},
      {tag:"2010"},
    ]
  },{
   id:2,
    tagType:"来源",
    tags:[
      {tag:"官方用语"},
      {tag:"国家事件"}
    ]
  }
])
const state=useNavigationStore()
const onRowExpand=function () {

}
onMounted(()=>{
  state.tagMode=true;
})
</script>
<style scoped>

</style>