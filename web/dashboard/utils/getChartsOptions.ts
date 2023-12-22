export default () => {
  // Composable
  const { isDarkMode } = useDarkMode()

  const chartBarOption = computed(() => {
    return {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        x: {
          ticks: {
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
          grid: {
            display: false,
          },
        },
        y: {
          ticks: {
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
          border: {
            display: false,
            dash: [2, 4],
          },
          grid: {
            lineWidth: 0.5,
            tickBorderDash: [2, 4],
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
        },
      },
      plugins: {
        legend: {
          display: false,
        },
        tooltip: {
          backgroundColor: isDarkMode.value ? '#272A2D' : undefined,
        },
      },
    }
  })

  const chartLineOption = computed(() => {
    return {
      responsive: true,
      maintainAspectRatio: false,
      tension: 0.4,
      interaction: {
        intersect: false,
        mode: 'nearest',
      },
      scales: {
        x: {
          border: {
            display: false,
            dash: [2, 4],
          },
          ticks: {
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
          grid: {
            lineWidth: 0.5,
            tickBorderDash: [2, 4],
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
        },
        y: {
          ticks: {
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
          border: {
            display: false,
            dash: [2, 4],
          },
          grid: {
            lineWidth: 0.5,
            tickBorderDash: [2, 4],
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
        },
      },
      plugins: {
        legend: {
          labels: {
            color: isDarkMode.value ? '#F9F9FB' : '#272A2D',
          },
        },
        tooltip: {
          backgroundColor: isDarkMode.value ? '#272A2D' : undefined,
        },
      },
    }
  })

  return { chartBarOption, chartLineOption }
}
