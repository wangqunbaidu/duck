<script lang="ts">
export default {
  setup() {
    return {}
  },
}
</script>

<template>
  <div>
    <MyHeader />

    <navs />

    <slot />
    <MyFooter />
  </div>
</template>

<style lang="scss" scoped>

</style>
