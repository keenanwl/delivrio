import {Component, EventEmitter, Input, Output} from '@angular/core';
import {TimeRangeSelectorModule} from "../time-range-selector/time-range-selector.module";
import {
	DateRange,
	DefaultMatCalendarRangeStrategy,
	MatDatepickerModule
} from "@angular/material/datepicker";
import {MatFormFieldModule} from "@angular/material/form-field";
import {DateTime} from "luxon";
import {provideLuxonDateAdapter} from "@angular/material-luxon-adapter";
import {TimeRange} from "../time-range-selector/time-range-selector.component";
import {MatButton} from "@angular/material/button";
import {MatIcon} from "@angular/material/icon";
import {NgIf} from "@angular/common";
import {MatDialogRef} from "@angular/material/dialog";

@Component({
	selector: 'app-date-time',
	standalone: true,
	imports: [
		TimeRangeSelectorModule,
		MatDatepickerModule,
		MatFormFieldModule,
		MatButton,
		MatIcon,
		NgIf,
	],
	providers: [provideLuxonDateAdapter(), DefaultMatCalendarRangeStrategy],
	templateUrl: './date-time.component.html',
	styleUrl: './date-time.component.scss'
})
export class DateTimeComponent {
	// Material messes with the time on update, so these are 2
	// independent vars until output
	end = DateTime.now().set({hour: 12, minute: 0});
	start = this.end.minus({day: 3});

	endTime = DateTime.fromISO("2023-01-01T00:00:00Z").set({hour: 16, minute: 0});
	startTime = this.endTime.set({hour: 7});

	@Input()
	get dateRange(): DateRange<DateTime> {
		return this._dateRange;
	}

	set dateRange(value: DateRange<DateTime>) {
		this._dateRange = value;
		this.start = this.start.set({
			year: this._dateRange.start?.year,
			month: this._dateRange.start?.month,
			day: this._dateRange.start?.day,
		});

		this.end = this.end.set({
			year: this._dateRange.end?.year,
			month: this._dateRange.end?.month,
			day: this._dateRange.end?.day,
		});

		this.startTime = this.startTime.set({
			hour: this._dateRange.start?.hour,
			minute: this._dateRange.start?.minute,
		});

		this.endTime = this.endTime.set({
			hour: this._dateRange.end?.hour,
			minute: this._dateRange.end?.minute,
		});
	}
	private _dateRange = new DateRange<DateTime>( this.start, this.end);

	@Output() dateRangeSelected = new EventEmitter<DateRange<DateTime>>();

	constructor(
		private readonly selectionStrategy: DefaultMatCalendarRangeStrategy<DateTime>,
		private dialogRef: MatDialogRef<any>
	) {
	}

	changed(d: DateTime) {
		this.dateRange = this.selectionStrategy.selectionFinished(d, this.dateRange)
	}

	timeChanged(times: TimeRange) {
		this.startTime = this.startTime.set({hour: times.start.hour, minute: times.start.minute});
		this.endTime = this.endTime.set({hour: times.end.hour, minute: times.end.minute});
	}

	publish() {
		this.dateRangeSelected.next(new DateRange<DateTime>(
			this.dateRange.start!.set({hour: this.startTime.hour, minute: this.startTime.minute}),
			this.dateRange.end!.set({hour: this.endTime.hour, minute: this.endTime.minute}),
		));
		this.close();
	}

	close() {
		this.dialogRef.close();
	}
}
