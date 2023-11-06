<script setup lang="ts">
import ListItem from "./ListItem.vue";
import ListHead from "./ListHead.vue";
import TagListItem from "./TagListItem.vue";
import {reactive} from "vue";
import {useNavigationStore} from "../state.ts";
const state=useNavigationStore()
const wordList=reactive([{
  word:"十七大",
  tags:["2007","国家事件","官方用语","直接引用","无衍生","正向","社会","社会生活"],
  comment:"十七大一般指中国共产党第十七次全国代表大会。"
},{
  word:"嫦娥一号",
  tags:["2007","国家事件","官方用语","直接引用","无衍生","正向","社会","社会生活"],
  comment:"嫦娥一号是中国探月计划中的第一颗绕月人造卫星，以中国古代神话人物嫦娥命名。"
},{
  word:"民生",
  tags:["2007","官方用语","直接引用","无衍生","中性","社会","社会生活"],
  comment:"民生，民是指人民，民生指人民的日常生活事项，例如衣、食、住、行、就业、娱乐、家庭、社团、公司、旅游等。广义上的民生概念是指，凡是同民生有关的，包括直接相关和间接相关的事情都属于民生范围内的事情。"
}])
</script>

<template>
  <div class="container mx-auto p-4 ">
    <h1 v-if="!state.tagMode" class="text-2xl font-bold mb-4">流行语</h1>
    <h1 v-if="state.tagMode" class="text-2xl font-bold mb-4">标签</h1>
    <ListHead/>
    <ul v-if="!state.tagMode" class="space-y-4 ">
      <ListItem v-for="word in wordList" :word="word.word" :tags="word.tags" :comment="word.comment"/>
    </ul>
  </div>
</template>

<style scoped>

</style>