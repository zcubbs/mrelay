<script setup lang="ts"></script>

<template>
  <div class="flex h-screen flex-col p-6">
    <!--  SIDEBAR  -->
    <ui-the-navbar />
    <!--  /SIDEBAR  -->

    <!--  PAGE  -->
    <div class="flex h-full flex-col py-6">
      <slot />
    </div>
    <!--  /PAGE  -->

    <!--  FOOTER  -->
    <ui-the-footer />
    <!--  /FOOTER  -->
  </div>
</template>
