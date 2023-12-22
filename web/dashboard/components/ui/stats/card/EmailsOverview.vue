<script setup lang="ts">
import { Bar } from 'vue-chartjs'

// Ref
const chartBar = ref()

// Utils
const { chartBarOption } = getChartsOptions()

const { chartData } = createChartData(
  chartBar,
  ['January', 'February', 'March', 'April', 'May'],
  [
    {
      label: 'Email',
      data: [40, 20, 12, 50, 10],
    },
  ]
)
</script>

<template>
  <shad-card class="border-primary bg-muted shadow dark:border-border dark:bg-background">
    <shad-card-header>
      <h4 class="text-lg">{{ $t('stats.card.overviewEmails.title') }}</h4>
    </shad-card-header>
    <shad-card-content class="h-96">
      <Bar ref="chartBar" :data="chartData" :options="chartBarOption" />
    </shad-card-content>
  </shad-card>
</template>
