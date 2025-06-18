import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { Job } from '../../models/job';

@Component({
  selector: 'app-job-card',
  standalone: true,
  imports: [CommonModule, MatCardModule],
  templateUrl: './job-card.component.html',
  styleUrls: ['./job-card.component.css']
})
export class JobCardComponent {
  @Input() job!: Job;
}
