<script setup lang="ts"></script>

<template>
  <div class="flex h-screen flex-col px-6 py-6">
    <!--  SIDEBAR  -->
    <ui-the-navbar />
    <!--  /SIDEBAR  -->

    <!--  PAGE  -->
    <slot />
    <!--  /PAGE  -->
  </div>
</template>
