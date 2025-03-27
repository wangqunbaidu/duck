<script setup>
useHead({
  title: "友情链接",
});
const links = await useHttpGet("/api/link/list");
</script>

<template>
  <section class="main">

  </section>
</template>

<style lang="scss" scoped></style>
