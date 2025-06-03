import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CommonModule } from '@angular/common';
import { MatDividerModule } from '@angular/material/divider';
import { JobCardComponent } from './components/job-card/job-card.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, MatDividerModule, JobCardComponent],
  templateUrl: './app.html',
  styleUrls: ['./app.css']
})
export class App {
  protected title = 'fullstackjobs-app';

  protected jobs = [
    {
      title: 'Fullstack Engineer',
      companyLogoUrl: 'https://logo.clearbit.com/google.com',
      country: 'Germany',
      city: 'Berlin',
      description: 'Join a fast-paced team building scalable systems.',
      link: '/job'
    },
    {
      title: 'Frontend Developer',
      companyLogoUrl: 'https://logo.clearbit.com/microsoft.com',
      country: 'Netherlands',
      city: 'Amsterdam',
      description: 'Work with modern Angular and reactive programming.',
      link: 'https://example.com/job2'
    },
    {
      title: 'Backend Developer',
      companyLogoUrl: 'https://logo.clearbit.com/facebook.com',
      country: 'United Kingdom',
      city: 'London',
      description: 'Develop RESTful APIs and microservices.',
      link: 'https://example.com/job3'
    },
    {
      title: 'DevOps Engineer',
      companyLogoUrl: 'https://logo.clearbit.com/amazon.com',
      country: 'Netherlands',
      city: 'Amsterdam',
      description: 'Automate deployment and manage cloud infrastructure.',
      link: 'https://example.com/job4'
    }
  ];
}
