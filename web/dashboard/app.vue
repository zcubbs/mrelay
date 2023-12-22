<script setup lang="ts">
import Toaster from '@/components/shad/toast/Toaster.vue'

// Composable
const { isDarkMode } = useDarkMode()

useHead({ title: 'Mail Relay', htmlAttrs: { class: computed(() => (isDarkMode.value ? 'dark' : '')) } })
</script>

<template>
  <NuxtLayout>
    <client-only>
      <Toaster />
    </client-only>
    <NuxtPage />
  </NuxtLayout>
</template>

<style>
.page-enter-active,
.page-leave-active {
  @apply duration-200;
}

.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style>
