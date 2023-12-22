<script setup lang="ts">
import { Line } from 'vue-chartjs'

// Ref
const chartLine = ref()

// Utils
const { chartLineOption } = getChartsOptions()

const { chartData } = createChartData(
  chartLine,
  ['January', 'February', 'March', 'April', 'May'],
  [
    {
      label: 'Alerts BL.Predict',
      data: [40, 20, 12, 50, 10],
    },
  ]
)
</script>

<template>
  <shad-card class="border-primary bg-muted shadow dark:border-border dark:bg-background">
    <shad-card-header>
      <h4 class="text-lg">{{ $t('stats.card.countEmails.title') }}</h4>
    </shad-card-header>
    <shad-card-content class="h-96">
      <Line ref="chartLine" :data="chartData" :options="chartLineOption" />
    </shad-card-content>
  </shad-card>
</template>
