import {Component, OnInit, ViewChild} from '@angular/core';
import {Actions, ofActionCompleted, Store} from '@ngxs/store';
import {DashboardActions} from "./dashboard.actions";
import FetchTiles = DashboardActions.FetchTiles;
import {Observable} from "rxjs";
import {DashboardModel, DashboardState} from "./dashboard.ngxs";
import {Paths} from "../../app-routing.module";
import {AppModel, AppState} from "../../app.ngxs";
import {
	ApexAxisChartSeries,
	ApexChart,
	ApexDataLabels,
	ApexGrid,
	ApexTitleSubtitle, ApexTooltip,
	ApexXAxis,
	ApexYAxis,
	ChartComponent
} from "ng-apexcharts";
import SetRateRequests = DashboardActions.SetRateRequests;
import {DatePipe} from "@angular/common";

export type ChartOptions = {
	series: ApexAxisChartSeries;
	chart: ApexChart;
	xaxis: ApexXAxis;
	yaxis: ApexYAxis;
	title: ApexTitleSubtitle;
	grid: ApexGrid,
	dataLabels: ApexDataLabels,
	tooltip: ApexTooltip,
};

@Component({
	selector: 'app-dashboard',
	templateUrl: './dashboard.component.html',
	styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

	app$: Observable<AppModel>;
	paths = Paths;
	state$: Observable<DashboardModel>;

	@ViewChild("chart") chart: ChartComponent | null = null;
	chartOptions: Partial<ChartOptions>;
	@ViewChild("chart-lookup") chartLookup: ChartComponent | null = null;
	chartOptionsLookup: Partial<ChartOptions>;

	constructor(private store: Store, private actions$: Actions, private datePipe: DatePipe) {
		this.app$ = store.select(AppState.get);
		this.state$ = store.select(DashboardState.get);

		const endX = new Date();
		// Set the start time to 6 hours ago from now
		const startX = new Date(endX.getTime() - (6 * 60 * 60 * 1000));

		this.chartOptionsLookup = {
			series: [{
					name: 'Options shown',
					data: []
				},
			],
			chart: {
				height: 350,
				type: 'scatter',
				zoom: {
					type: 'xy',
					enabled: true,
					autoScaleYaxis: true
				}
			},
			dataLabels: {
				enabled: false
			},
			grid: {
				xaxis: {
					lines: {
						show: true
					}
				},
				yaxis: {
					lines: {
						show: true
					}
				},
			},
			xaxis: {
				type: 'datetime',
				min: startX.getTime(),
				max: endX.getTime(),
				labels: {
					datetimeUTC: false,
				},
			},
			yaxis: {
				tickAmount: 10,
				stepSize: 1,
				min: 0,
			},
			tooltip: {
				custom: ({series, seriesIndex, dataPointIndex, w})=> {
					const dataPoint = w.globals.series[seriesIndex][dataPointIndex];
					const customData = w.globals.initialSeries[seriesIndex].data[dataPointIndex].meta;
					const dateFormatted = this.datePipe.transform(
						new Date(w.globals.seriesX[seriesIndex][dataPointIndex]),
						'yyyy-MM-dd HH:mm:ss'
					);
					return `
					  <div class="apexcharts-tooltip-title">
						${dateFormatted}
					  </div>
					  <div class="apexcharts-tooltip-y-group">
						<span class="apexcharts-tooltip-text-y-label">Options shown: </span>
						<span class="apexcharts-tooltip-text-y-value">${dataPoint}</span>
					  </div>
					  <div>${customData}</div>
					`;

				}
			}
		};
		this.chartOptions = {
			series: [],
			chart: {
				height: 350,
				type: "bar",
				toolbar: {show: false},
			},
			title: {
				text: "Count of products last updated by day"
			},
			xaxis: {
				categories: this.getLast7Days(),
			}
		};
	}

	ngOnInit() {
		this.store.dispatch(new FetchTiles());
		this.actions$.pipe(ofActionCompleted(SetRateRequests))
			.subscribe((p) => {
				this.chartOptionsLookup.series = [
					{
						name: "Options shown",
						data: p.action.payload.requests.map((r) => {
							return {
								x: new Date(r.date).getTime(),
								y: r.optionCount,
								meta: `
									${JSON.stringify(JSON.parse(r.req || '{}'), null, 2)}
								`,
							}
						})
					}, {
						name: "Errors",
						data: p.action.payload.requestsError.map((r) => {
							return {
								x: new Date(r.date).getTime(),
								y: r.optionCount,
								meta: `
									${r.error}
									${JSON.stringify(JSON.parse(r.req || '{}'), null, 2)}
								`,
							}
						})
					}
				];
			});
	}

	getLast7Days() {
		const daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
		const last7Days = [];

		for (let i = 6; i >= 0; i--) {
			const date = new Date();
			date.setDate(date.getDate() - i);
			if (i === 0) {
				last7Days.push('Today');
			} else if (i === 1) {
				last7Days.push('Yesterday');
			} else {
				last7Days.push(daysOfWeek[date.getDay()]);
			}
		}

		return last7Days;
	}

}
