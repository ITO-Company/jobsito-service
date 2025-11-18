package kpi

/*
KPI MODULE - COMPLETE OVERVIEW

This module provides comprehensive KPI tracking for your intership management system.
All KPIs are context-aware (company-level or intership-level).

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. MILESTONE KPI (milestone_kpi_*.go)
   ├─ CompletionRate (%)
   ├─ AverageCompletionTime (days)
   ├─ OverduePercentage (%)
   ├─ Status breakdown: Completed, Pending, Active, Overdue
   └─ Contexts: Company-level, Intership-level

   Endpoints:
   - GET /kpis/milestones/company (requires company role)
   - GET /kpis/milestones/intership/:intership_id

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

2. ISSUE KPI (issue_kpi_*.go)
   ├─ ResolutionRate (%)
   ├─ AverageResolutionTime (days)
   ├─ OverduePercentage (%)
   ├─ IssuesWithRequests (count)
   ├─ Status breakdown: Resolved, Pending, Active, Overdue
   └─ Contexts: Company-level, Intership-level

   Endpoints:
   - GET /kpis/issues/company (requires company role)
   - GET /kpis/issues/intership/:intership_id

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

3. REQUEST KPI (request_kpi_*.go)
   ├─ ApprovalRate (%)
   ├─ RejectionRate (%)
   ├─ AverageReviewTime (days)
   ├─ PendingPercentage (%)
   ├─ IssuesWithPendingRequests (count)
   ├─ Status breakdown: Approved, Rejected, Pending
   └─ Contexts: Company-level, Intership-level

   Endpoints:
   - GET /kpis/requests/company (requires company role)
   - GET /kpis/requests/intership/:intership_id

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

4. CONVERSION KPI (conversion_kpi_*.go)
   ├─ ApplicationAcceptanceRate (%)
   ├─ ConversionToIntershipRate (%)
   ├─ IntershipCompletionRate (%)
   ├─ AvgTimeAppToAcceptance (days)
   ├─ AvgTimeAcceptanceToIntership (days)
   ├─ SalaryAnalysis: ProposedAvg vs OfferedMin/Max
   ├─ Funnel: TotalApps → AcceptedApps → InitiatedInterships → CompletedInterships
   └─ Contexts: Company-level, Job Posting-level

   Endpoints:
   - GET /kpis/conversions/company (requires company role)
   - GET /kpis/conversions/job-posting/:job_posting_id

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

ARCHITECTURE:

Handler (HTTP) → Service (Business Logic) → Repo (Database Queries)
   ↓                    ↓                      ↓
kpi_handler.go    milestone_kpi_service.go  milestone_kpi_repo.go
                  issue_kpi_service.go      issue_kpi_repo.go
                  request_kpi_service.go    request_kpi_repo.go
                  conversion_kpi_service.go conversion_kpi_repo.go

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

EXAMPLE RESPONSES:

1. Milestone KPI Response:
{
  "company_id": "123e4567-e89b-12d3-a456-426614174000",
  "total_milestones": 10,
  "completed_milestones": 8,
  "pending_milestones": 1,
  "active_milestones": 1,
  "overdue_milestones": 0,
  "completion_rate": 80.0,
  "average_completion_time_days": 15.5,
  "overdue_percentage": 0.0
}

2. Conversion KPI Response:
{
  "context_id": "comp-123",
  "context_type": "company",
  "total_applications": 50,
  "accepted_applications": 20,
  "initiated_interships": 18,
  "completed_interships": 15,
  "application_acceptance_rate": 40.0,
  "conversion_to_intership_rate": 90.0,
  "intership_completion_rate": 83.33,
  "avg_time_app_to_acceptance_days": 3.5,
  "avg_time_acceptance_to_intership_days": 2.1,
  "avg_proposed_salary": "4500",
  "avg_offered_salary_min": "3000",
  "avg_offered_salary_max": "6000"
}

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

KEY INSIGHTS FROM KPIs:

📊 Milestone Health:
- If completion_rate < 70% → Gestión deficiente
- If average_completion_time > 30 days → Proyecto lento
- If overdue_percentage > 20% → Problemas de seguimiento

📊 Issue Resolution:
- If resolution_rate < 80% → Muchos problemas sin resolver
- If average_resolution_time > 7 days → Lento en resolver
- If issues_with_requests > 50% → Muchos desacuerdos

📊 Request Management:
- If approval_rate < 70% → Empresa muy restrictiva
- If pending_percentage > 20% → Lento en revisiones
- If average_review_time > 5 days → Respuestas lentas

📊 Conversion Funnel:
- If application_acceptance_rate < 20% → Ofertas poco atractivas
- If conversion_to_intership_rate < 50% → Candidatos se retractan
- If intership_completion_rate < 70% → Malas pasantías
- Salary mismatch → Si proposed < offered_min × 0.8 → Candidato pierde interés

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
*/
