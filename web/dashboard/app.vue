<script setup lang="ts">
// Composable
const { isDarkMode } = useDarkMode()

useHead({ title: 'Mail Relay', htmlAttrs: { class: computed(() => (isDarkMode.value ? 'dark' : '')) } })
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
