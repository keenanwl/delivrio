import {Pipe, PipeTransform} from '@angular/core';
import {WorkstationStatus} from "../../../../generated/graphql";

@Pipe({
	name: 'workstationStatusPipe',
	standalone: true
})
export class WorkstationStatusPipePipe implements PipeTransform {

  transform(value: WorkstationStatus): string {
    switch (value) {
		case WorkstationStatus.Active:
			return "#4fb500";
		case WorkstationStatus.Pending:
			return "#cb011c";
		case WorkstationStatus.Disabled:
			return "#000000";
		case WorkstationStatus.Offline:
			return "#595959";
	}
  }

}
