(() => {
  Array.from(document.querySelectorAll("form.unreproduced")).map(form => {

    const chartDiv = document.createElement("div");
    chartDiv.setAttribute("class", "drv-pie-chart")

    form.insertBefore(chartDiv, form.firstChild)

    const columns = []

    Array.from(form.querySelectorAll("div.drv-output")).map(output => {
      const outputHash = output.attributes["x-data-outputhash"].value
      const numLogs = output.querySelectorAll("ul.log-names > li.log-name").length
      columns.push([outputHash, numLogs])
    })

    const chart = c3.generate({
      bindto: chartDiv,
      padding: {
        left: 0,
      },
      data: {
        type : 'pie',
        onclick: (d, i) => {
          form.querySelector(`input[type=checkbox][name='output_hash'][value='${d.id}']`).click()
        },
        columns: columns,
      }
    });

  })
})()