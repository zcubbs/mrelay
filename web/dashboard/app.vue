<script setup lang="ts">
// Composable
const { isDarkMode } = useDarkMode()

useHead({ htmlAttrs: { class: computed(() => (isDarkMode.value ? 'dark' : '')) } })
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
