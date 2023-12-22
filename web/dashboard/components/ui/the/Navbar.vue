<script setup lang="ts">
import { addDays } from 'date-fns'
import { breakpointsTailwind } from '@vueuse/core'

// Composable
const route = useRoute()
const { formatTime } = getTime()
const { isDarkMode, setDarkMode } = useDarkMode()
const breakpoints = useBreakpoints(breakpointsTailwind)

// Ref
const date = ref({
  start: addDays(new Date(), -7),
  end: new Date(),
})

// Computed
const isStatsPage = computed(() => route.path === '/')

const menu = computed(() => {
  return {
    path: isStatsPage.value ? '/form' : '/',
    identifier: isStatsPage.value ? 'navbar.menu.form' : 'navbar.menu.home',
  }
})

const titleIdentifier = computed(() => {
  return isStatsPage.value ? 'navbar.title.stats' : 'navbar.title.form'
})

const alignCalendar = computed(() => {
  return breakpoints.greaterOrEqual('sm').value ? 'end' : 'center'
})

// Methods
const toggleDarkMode = () => {
  setDarkMode(!isDarkMode.value)
}
</script>

<template>
  <div class="flex flex-col items-center justify-between gap-y-5 sm:flex-row sm:items-start">
    <div class="flex items-baseline gap-x-3">
      <h1 class="text-2xl font-bold">Mail Relay</h1>
      <span class="text-xl font-bold">/</span>
      <NuxtLink :to="menu.path" class="group transition duration-300">
        <span>{{ $t(menu.identifier) }}</span>
        <span class="block h-0.5 max-w-0 bg-foreground transition-all duration-500 group-hover:max-w-full"></span>
      </NuxtLink>
    </div>

    <div class="hidden items-center gap-x-3 lg:flex">
      <PhosphorIconChartLine v-if="isStatsPage" size="20" />
      <PhosphorIconTextbox v-else size="20" />
      <h2 class="text-xl">{{ $t(titleIdentifier) }}</h2>
    </div>

    <div class="flex items-center gap-x-5">
      <shad-popover v-if="isStatsPage">
        <shad-popover-trigger as-child>
          <shad-button id="date">
            <PhosphorIconCalendarBlank size="16" class="mr-2" />
            <span> {{ formatTime(date.start.toString(), 'DD/MM/YYYY') }} - {{ formatTime(date.end.toString(), 'DD/MM/YYYY') }} </span>
          </shad-button>
        </shad-popover-trigger>
        <shad-popover-content class="w-auto p-0" :align="alignCalendar" :avoid-collisions="true">
          <shad-calendar v-model.range="date" :columns="2" />
        </shad-popover-content>
      </shad-popover>

      <div class="cursor-pointer rounded p-1.5 duration-200 hover:bg-primary/10" @click="toggleDarkMode">
        <PhosphorIconSun v-if="isDarkMode" size="20" />
        <PhosphorIconMoon v-else size="20" />
      </div>
    </div>
  </div>
</template>
