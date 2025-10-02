import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';

export type TimeRange = {start: {hour: number; minute: number;}; end: {hour: number; minute: number;}};

@Component({
	selector: 'app-time-range-selector',
	templateUrl: './time-range-selector.component.html',
	styleUrls: ['./time-range-selector.component.scss']
})
export class TimeRangeSelectorComponent implements OnInit {

	@Input() hourFrom = 9;
	@Input() minutesFrom = 0;

	@Input() hourTo = 15;
	@Input() minutesTo = 0;

	@Input() hideDialog = false;

	invalid = false;

	@Output() selectedRange = new EventEmitter<{start: string; end: string}>();
	@Output() selectedRangeRaw = new EventEmitter<TimeRange>();

	constructor(private dialogRef: MatDialogRef<TimeRangeSelectorComponent>) { }

	ngOnInit(): void {
	}

	incrementHourFrom() {
		this.hourFrom = this.next(this.hourFrom, 24, 1);
		this.invalid = this.isInvalid();
		this.publish();
	}

	incrementHourTo() {
		this.hourTo = this.next(this.hourTo, 24, 1);
		this.invalid = this.isInvalid();
		this.publish();
	}

	incrementMinutesFrom() {
		this.minutesFrom = this.next(this.minutesFrom, 60, 5);
		this.invalid = this.isInvalid();
		this.publish();
	}

	incrementMinutesTo() {
		this.minutesTo = this.next(this.minutesTo, 60, 5);
		this.invalid = this.isInvalid();
		this.publish();
	}

	decrementHoursFrom() {
		this.hourFrom = this.previous(this.hourFrom, 1);
		this.invalid = this.isInvalid();
		this.publish();
	}

	decrementHoursTo() {
		this.hourTo = this.previous(this.hourTo, 1);
		this.invalid = this.isInvalid();
		this.publish();
	}

	decrementMinutesFrom() {
		this.minutesFrom = this.previous(this.minutesFrom, 5);
		this.invalid = this.isInvalid();
		this.publish();
	}

	decrementMinutesTo() {
		this.minutesTo = this.previous(this.minutesTo, 5);
		this.invalid = this.isInvalid();
		this.publish();
	}

	isInvalid(): boolean {
		if (this.minutesFrom >= this.minutesTo && this.hourFrom >= this.hourTo) {
			return true;
		}
		return false;
	}

	next(current: number, max: number, step: number): number {
		if (current + step >= 59) {
			return 59
		} else if (current + step < max) {
			return current + step;
		}
		return current;
	}

	previous(current: number, step: number): number {
		if (current === 59) {
			return 60 - step
		} else if (current - step > -1) {
			return current - step;
		}
		return current;
	}

	close() {
		this.dialogRef.close();
	}

	publish() {
		this.selectedRange.next({
			start: `${String(this.hourFrom).padStart(2, '0')}:${String(this.minutesFrom).padStart(2, '0')}`,
			end: `${String(this.hourTo).padStart(2, '0')}:${String(this.minutesTo).padStart(2, '0')}`,
		});
		this.selectedRangeRaw.next({
			start: {hour: this.hourFrom, minute: this.minutesFrom},
			end: {hour: this.hourTo, minute: this.minutesTo},
		})
	}

	save() {
		this.publish();
		this.close();
	}

}
