import type { Ref } from 'vue'

export default (chartRef: Ref, labels: string[], datasets: { label: string; data: any[] }[]) => {
  // Computed property for chart data
  const chartData = computed(() => {
    return {
      labels,
      datasets: datasets.map((dataset) => ({
        label: dataset.label,
        data: dataset.data,
        borderRadius: 4,
      })),
    }
  })

  // Function to add data to the chart
  const addData = (label: string, data: any[]) => {
    chartData.value.labels.push(label)

    data.forEach((d, idx) => chartData.value.datasets[idx].data.push(d))

    chartRef.value.chart.update()
  }

  // Function to remove data from the chart
  const removeData = (dataIndex: number) => {
    if (dataIndex < 0) {
      alert(`Invalid dataIndex: ${dataIndex}`)
      return
    }

    chartData.value.labels.splice(dataIndex, 1)

    chartData.value.datasets.forEach((dataset) => {
      if (dataIndex < dataset.data.length) {
        dataset.data.splice(dataIndex, 1)
      }
    })

    chartRef.value.chart.update()
  }

  return { chartData, addData, removeData }
}
