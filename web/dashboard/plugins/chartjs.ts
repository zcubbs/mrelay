import { Chart, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement, Colors } from 'chart.js'

export default defineNuxtPlugin(() => {
  Chart.register(Colors, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend, PointElement, LineElement)
})
