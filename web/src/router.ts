import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component:()=>import("./components/List.vue")},
        { path:"/tag",component:()=>import("./components/TagList.vue")}
    ]
});