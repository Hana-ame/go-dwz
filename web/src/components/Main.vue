<script setup lang="ts">
import LinkContainer from "@/components/LinkContainer.vue";


const tags = ["public" , "private"];
// const tags = ["tag1"];

</script>

<template>
  <main>
    <LinkContainer v-for="tag in tags" :key="tag" :tag="tag"></LinkContainer>
  </main>
</template>

<style>

main {
  padding: 1em;
}

</style>
