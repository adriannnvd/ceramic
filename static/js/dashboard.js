const xValues = ["April", "May", "June", "July", "August"];
const good = [25, 30, 35, 40, 45]; // First set of data
const defect = [30, 19, 9, 14, 8];   // Second set of data
const barColor1 = "green";
const barColor2 = "red";


new Chart("myChart", {
  type: "bar",
  data: {
    labels: xValues,
    datasets: [
      {
        label: 'Good', // Label for the first set of bars
        backgroundColor: barColor1,
        data: good
      },
      {
        label: 'Defect', // Label for the second set of bars
        backgroundColor: barColor2, // Use the same colors as the first set
        data: defect
      }
    ]
  },
  options: {
    legend: {
      display: true,
      position: 'bottom' // Place legend at the bottom
    },
    title: {
      display: true,
      text: "Ceramic Plate Classification Uploads"
    }
  }
});
