<script setup lang="ts">
import Banner from "@/components/Banner.vue";
import Main from "@/components/Main.vue";
</script>

<template>
  <Banner />
  <Main />
</template>

<style>
body {
  background-attachment: fixed;
  /* background-image: url("https://api.ixiaowai.cn/api/api.php"); */
  background-color: black;
  background-repeat: no-repeat;
  background-size: cover;
  background-position-x: center;
  background-position-y: center;
  overflow-y: scroll;
}


main {
  background-color: rgba(255, 255, 255, 0.5);
}

.s-title {
  font-size: 1.5em;
  font-weight: 700;
  color: black;
  display: block;
}

h1 {
  text-shadow: 0.2rem 0rem 0.5rem white, -0.2rem 0rem 0.5rem white,
    0rem 0.2rem 0.5rem white, 0rem -0.2rem 0.5rem white;
}

</style>
