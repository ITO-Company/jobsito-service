-- ============================================================================
-- SEED DATA MASIVO PARA JOBSITO - VERSIÓN FINAL Y FUNCIONAL
-- ============================================================================

-- ============================================================================
-- 1. GLOBAL TAGS (100 tags relacionados con TI)
-- ============================================================================
INSERT INTO global_tags (id, name, category, color, is_approved, usage_count, created_at, updated_at)
VALUES
(gen_random_uuid(), 'Python', 'Programming Language', '#3776AB', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'JavaScript', 'Programming Language', '#F7DF1E', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Java', 'Programming Language', '#007396', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'C++', 'Programming Language', '#00599C', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'C#', 'Programming Language', '#239120', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Go', 'Programming Language', '#00ADD8', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Rust', 'Programming Language', '#CE422B', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'PHP', 'Programming Language', '#777BB4', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Ruby', 'Programming Language', '#CC342D', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'TypeScript', 'Programming Language', '#3178C6', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'React', 'Frontend Framework', '#61DAFB', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Vue.js', 'Frontend Framework', '#4FC08D', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Angular', 'Frontend Framework', '#DD0031', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Svelte', 'Frontend Framework', '#FF3E00', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Next.js', 'Frontend Framework', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Express.js', 'Backend Framework', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Django', 'Backend Framework', '#092E20', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Flask', 'Backend Framework', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Spring Boot', 'Backend Framework', '#6DB33F', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'FastAPI', 'Backend Framework', '#009688', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Node.js', 'Runtime', '#68A063', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Docker', 'DevOps', '#2496ED', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Kubernetes', 'DevOps', '#326CE5', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'AWS', 'Cloud', '#FF9900', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Azure', 'Cloud', '#0078D4', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Google Cloud', 'Cloud', '#4285F4', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'PostgreSQL', 'Database', '#336791', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'MySQL', 'Database', '#00758F', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'MongoDB', 'Database', '#13AA52', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Redis', 'Database', '#DC382D', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Elasticsearch', 'Search Engine', '#005571', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Git', 'Version Control', '#F05032', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'GitHub', 'Version Control', '#181717', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'GitLab', 'Version Control', '#FCA121', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Jenkins', 'CI/CD', '#D24939', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'GitLab CI', 'CI/CD', '#FCA121', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'GitHub Actions', 'CI/CD', '#2088FF', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Terraform', 'IaC', '#7B42BC', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Ansible', 'Configuration', '#EE0000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Linux', 'Operating System', '#FCC624', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Windows Server', 'Operating System', '#0078D4', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'macOS', 'Operating System', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'REST API', 'Architecture', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'GraphQL', 'Query Language', '#E10098', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'gRPC', 'Communication', '#244C5A', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Microservices', 'Architecture', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Monolithic', 'Architecture', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Machine Learning', 'AI', '#FF6B00', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Deep Learning', 'AI', '#0078D4', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'TensorFlow', 'ML Framework', '#FF6F00', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'PyTorch', 'ML Framework', '#EE4C2C', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Scikit-learn', 'ML Library', '#F7931E', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Pandas', 'Data Science', '#150458', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'NumPy', 'Data Science', '#013243', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Jupyter', 'Notebook', '#F37726', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Spark', 'Big Data', '#E25A1C', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Hadoop', 'Big Data', '#FFCC00', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Apache Kafka', 'Message Queue', '#231F20', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'RabbitMQ', 'Message Queue', '#FF6600', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Vue', 'Frontend Framework', '#4FC08D', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'HTML5', 'Markup', '#E34C26', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'CSS3', 'Styling', '#1572B6', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Bootstrap', 'CSS Framework', '#7952B3', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Tailwind CSS', 'CSS Framework', '#06B6D4', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Material Design', 'Design System', '#757575', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Webpack', 'Bundler', '#8DD6F9', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Vite', 'Build Tool', '#646CFF', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Jest', 'Testing', '#15C213', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Mocha', 'Testing', '#8D6748', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Pytest', 'Testing', '#0A9EDC', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Selenium', 'Testing', '#43B02A', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Cypress', 'Testing', '#17202C', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'QA Testing', 'Quality Assurance', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Security', 'Cybersecurity', '#FF0000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'OAuth 2.0', 'Authentication', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'JWT', 'Authentication', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'API Gateway', 'Networking', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Load Balancer', 'Networking', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Nginx', 'Web Server', '#009639', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Apache', 'Web Server', '#D70015', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'SSL/TLS', 'Security', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Agile', 'Methodology', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Scrum', 'Methodology', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Kanban', 'Methodology', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Jira', 'Project Management', '#0052CC', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Trello', 'Project Management', '#0079BF', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Slack', 'Communication', '#E01E5A', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Figma', 'Design', '#F24E1E', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Adobe XD', 'Design', '#FF61F6', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'UX/UI Design', 'Design', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'SEO', 'Marketing', '#4CAF50', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Analytics', 'Data Analysis', '#E37400', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Data Warehouse', 'Data Storage', '#336791', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'ETL', 'Data Integration', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Data Mining', 'Data Analysis', '#FF6B35', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Blockchain', 'Technology', '#F7931A', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Smart Contracts', 'Blockchain', '#3C3C3D', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Web3', 'Blockchain', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'IoT', 'Internet of Things', '#0096D6', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Arduino', 'Hardware', '#00979D', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Raspberry Pi', 'Hardware', '#C51A4A', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Android Development', 'Mobile', '#3DDC84', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'iOS Development', 'Mobile', '#000000', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'Flutter', 'Mobile Framework', '#02569B', 'true', '0', NOW(), NOW()),
(gen_random_uuid(), 'React Native', 'Mobile Framework', '#61DAFB', 'true', '0', NOW(), NOW());

-- ============================================================================
-- 2. COMPANY PROFILES (1000 empresas)
-- ============================================================================
INSERT INTO company_profiles (id, company_name, password, email, description, website, phone, address, industry, company_size, logo_url, is_verified, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    'Tech Corp ' || i,
    'hashed_password_' || i,
    'company' || i || '@techcorp.com',
    'We are a software development company specialized in enterprise solutions. Company number ' || i,
    'https://techcorp' || i || '.com',
    '+56-2-' || LPAD((i % 10000)::text, 8, '0'),
    'Avenida Tecnológica ' || i || ', Santiago, Chile',
    CASE WHEN i % 5 = 0 THEN 'Software Development'
         WHEN i % 5 = 1 THEN 'Cloud Services'
         WHEN i % 5 = 2 THEN 'AI/ML Solutions'
         WHEN i % 5 = 3 THEN 'Cybersecurity'
         ELSE 'Data Analytics' END,
    CASE WHEN i % 4 = 0 THEN '10-50'
         WHEN i % 4 = 1 THEN '50-200'
         WHEN i % 4 = 2 THEN '200-1000'
         ELSE '1000+' END,
    'https://logo.techcorp' || i || '.png',
    true,
    NOW(),
    NOW()
FROM generate_series(1, 1000) AS t(i);

-- ============================================================================
-- 3. JOB SEEKER PROFILES (1000 buscadores de trabajo)
-- ============================================================================
INSERT INTO job_seeker_profiles (id, name, password, email, bio, phone, location, cv_url, portfolio_url, expected_salary_min, expected_salary_max, availability, skills, experience, is_active, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    'Job Seeker ' || i,
    'hashed_password_' || i,
    'jobseeker' || i || '@email.com',
    'Passionate developer with expertise in software development. Professional #' || i,
    '+56-9-' || LPAD((i % 100000)::text, 8, '0'),
    CASE WHEN i % 16 = 0 THEN 'Santiago'
         WHEN i % 16 = 1 THEN 'Valparaíso'
         WHEN i % 16 = 2 THEN 'Concepción'
         WHEN i % 16 = 3 THEN 'La Serena'
         WHEN i % 16 = 4 THEN 'Valdivia'
         WHEN i % 16 = 5 THEN 'Puerto Montt'
         WHEN i % 16 = 6 THEN 'Coyhaique'
         WHEN i % 16 = 7 THEN 'Punta Arenas'
         WHEN i % 16 = 8 THEN 'Antofagasta'
         WHEN i % 16 = 9 THEN 'Calama'
         WHEN i % 16 = 10 THEN 'Iquique'
         WHEN i % 16 = 11 THEN 'Arica'
         WHEN i % 16 = 12 THEN 'Los Ángeles'
         WHEN i % 16 = 13 THEN 'Talca'
         WHEN i % 16 = 14 THEN 'Rancagua'
         ELSE 'Temuco' END,
    'https://cv.storage.com/cv_' || i || '.pdf',
    'https://portfolio' || i || '.com',
    '1500000',
    '3500000',
    CASE WHEN i % 3 = 0 THEN 'Full Time'
         WHEN i % 3 = 1 THEN 'Part Time'
         ELSE 'Contract' END,
    'Software Development, Web Development, Database Design',
    CASE WHEN i % 3 = 0 THEN '1-2 years'
         WHEN i % 3 = 1 THEN '3-5 years'
         ELSE '5+ years' END,
    true,
    NOW(),
    NOW()
FROM generate_series(1, 1000) AS t(i);

-- ============================================================================
-- 4. JOB POSTINGS (3000 ofertas de trabajo)
-- ============================================================================
INSERT INTO job_postings (id, title, description, requirement, salary_min, salary_max, work_type, experience_level, location, is_remote, is_hibrid, contract_type, benefit, status, is_closed, expires_at, company_profile_id, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    CASE WHEN (i % 10) = 0 THEN 'Senior Backend Developer'
         WHEN (i % 10) = 1 THEN 'Full Stack Developer'
         WHEN (i % 10) = 2 THEN 'Frontend Developer'
         WHEN (i % 10) = 3 THEN 'DevOps Engineer'
         WHEN (i % 10) = 4 THEN 'Data Scientist'
         WHEN (i % 10) = 5 THEN 'Mobile Developer'
         WHEN (i % 10) = 6 THEN 'QA Engineer'
         WHEN (i % 10) = 7 THEN 'Cloud Architect'
         WHEN (i % 10) = 8 THEN 'Security Engineer'
         ELSE 'Database Administrator' END || ' - ' || i,
    'We are looking for a talented professional to join our team. Position #' || i,
    '- Bachelor in Computer Science or equivalent\n- 2+ years of experience\n- Strong problem-solving skills',
    CASE WHEN (i % 5) = 0 THEN '3000000'
         WHEN (i % 5) = 1 THEN '2500000'
         WHEN (i % 5) = 2 THEN '4000000'
         WHEN (i % 5) = 3 THEN '3500000'
         ELSE '2000000' END,
    CASE WHEN (i % 5) = 0 THEN '5000000'
         WHEN (i % 5) = 1 THEN '4500000'
         WHEN (i % 5) = 2 THEN '6000000'
         WHEN (i % 5) = 3 THEN '5500000'
         ELSE '4000000' END,
    'Full Time',
    CASE WHEN (i % 4) = 0 THEN 'Junior (0-2 years)'
         WHEN (i % 4) = 1 THEN 'Mid-level (2-5 years)'
         WHEN (i % 4) = 2 THEN 'Senior (5+ years)'
         ELSE 'Lead' END,
    CASE WHEN i % 10 = 0 THEN 'Remote'
         WHEN i % 10 = 1 THEN 'Hybrid'
         ELSE 'On-site' END,
    i % 3 = 0,
    i % 3 = 1,
    'Full Time',
    'Health Insurance, Flexible Hours, Home Office, Training Budget',
    'Active',
    false,
    NOW() + INTERVAL '90 days',
    (SELECT id FROM company_profiles ORDER BY RANDOM() LIMIT 1),
    NOW(),
    NOW()
FROM generate_series(1, 3000) AS t(i);

-- ============================================================================
-- 5. JOB POSTING TAGS (0-50 tags aleatorios por job posting)
-- ============================================================================
INSERT INTO job_posting_tags (id, job_posting_id, global_tag_id, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    jp.id,
    (SELECT id FROM global_tags ORDER BY RANDOM() LIMIT 1),
    NOW(),
    NOW()
FROM job_postings jp,
     generate_series(1, FLOOR(RANDOM() * 50)::int) AS seq;

-- ============================================================================
-- 6. JOB SEEKER TAGS (0-50 tags aleatorios por job seeker)
-- ============================================================================
INSERT INTO job_seeker_tags (id, job_seeker_profile_id, global_tag_id, proficiency_level, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    jsp.id,
    (SELECT id FROM global_tags ORDER BY RANDOM() LIMIT 1),
    CASE WHEN RANDOM() < 0.3 THEN 'Beginner'
         WHEN RANDOM() < 0.7 THEN 'Intermediate'
         ELSE 'Advanced' END,
    NOW(),
    NOW()
FROM job_seeker_profiles jsp,
     generate_series(1, FLOOR(RANDOM() * 50)::int) AS seq;

-- ============================================================================
-- 7. INTERNSHIPS (100 pasantías)
-- ============================================================================
INSERT INTO interships (id, start_date, end_date, status, job_posting_id, job_seeker_profile_id, company_profile_id, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    NOW() - INTERVAL '60 days',
    NOW() + INTERVAL '60 days',
    CASE WHEN RANDOM() < 0.3 THEN 'pending'
         WHEN RANDOM() < 0.6 THEN 'active'
         WHEN RANDOM() < 0.8 THEN 'completed'
         ELSE 'rejected' END,
    (SELECT id FROM job_postings ORDER BY RANDOM() LIMIT 1),
    (SELECT id FROM job_seeker_profiles ORDER BY RANDOM() LIMIT 1),
    (SELECT id FROM company_profiles ORDER BY RANDOM() LIMIT 1),
    NOW(),
    NOW()
FROM generate_series(1, 100) AS t(i);

-- ============================================================================
-- 8. FOLLOWUP MILESTONES (2-5 por pasantía)
-- ============================================================================
INSERT INTO followup_milestones (id, title, description, due_date, status, intership_id, company_profile_id, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    'Milestone ' || seq,
    'Milestone description for internship',
    NOW() + (INTERVAL '1 day' * (seq * 15)),
    CASE WHEN RANDOM() < 0.3 THEN 'pending'
         WHEN RANDOM() < 0.6 THEN 'active'
         WHEN RANDOM() < 0.8 THEN 'completed'
         ELSE 'rejected' END,
    i.id,
    i.company_profile_id,
    NOW(),
    NOW()
FROM interships i,
     generate_series(1, FLOOR(RANDOM() * 4 + 2)::int) seq;

-- ============================================================================
-- 9. FOLLOWUP ISSUES (1-3 por milestone)
-- ============================================================================
INSERT INTO followup_issues (id, title, description, due_date, status, followup_milestone_id, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    'Issue: Task ' || seq,
    'Issue description for milestone',
    fm.due_date - INTERVAL '5 days' + (INTERVAL '1 day' * seq),
    CASE WHEN RANDOM() < 0.3 THEN 'pending'
         WHEN RANDOM() < 0.6 THEN 'active'
         WHEN RANDOM() < 0.8 THEN 'completed'
         ELSE 'rejected' END,
    fm.id,
    NOW(),
    NOW()
FROM followup_milestones fm,
     generate_series(1, FLOOR(RANDOM() * 3 + 1)::int) seq;

-- ============================================================================
-- 10. REQUESTS (1-2 por issue)
-- ============================================================================
INSERT INTO requests (id, title, description, status, company_comment, followup_issue_id, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    'Request ' || seq,
    'Request description for issue',
    CASE WHEN RANDOM() < 0.3 THEN 'pending'
         WHEN RANDOM() < 0.6 THEN 'approved'
         WHEN RANDOM() < 0.8 THEN 'rejected'
         ELSE 'in_review' END,
    CASE WHEN RANDOM() < 0.5 THEN 'Great work! Keep it up.'
         WHEN RANDOM() < 0.7 THEN 'Needs revisions. Please review the feedback.'
         ELSE 'Excellent execution.' END,
    fi.id,
    NOW(),
    NOW()
FROM followup_issues fi,
     generate_series(1, FLOOR(RANDOM() * 2 + 1)::int) seq;

-- ============================================================================
-- 11. SAVED JOBS (7500 aprox)
-- ============================================================================
INSERT INTO saved_jobs (id, job_seeker_profile_id, job_posting_id, created_at, updated_at)
SELECT DISTINCT ON (seeker_id, posting_id)
    gen_random_uuid(),
    seeker_id,
    posting_id,
    NOW(),
    NOW()
FROM (
    SELECT 
        (SELECT id FROM job_seeker_profiles ORDER BY RANDOM() LIMIT 1) as seeker_id,
        (SELECT id FROM job_postings ORDER BY RANDOM() LIMIT 1) as posting_id
    FROM generate_series(1, 7500)
) t;

-- ============================================================================
-- VERIFICATION
-- ============================================================================
SELECT 
    'Global Tags' AS table_name, COUNT(*) AS count FROM global_tags
UNION ALL
SELECT 'Company Profiles', COUNT(*) FROM company_profiles
UNION ALL
SELECT 'Job Postings', COUNT(*) FROM job_postings
UNION ALL
SELECT 'Job Seeker Profiles', COUNT(*) FROM job_seeker_profiles
UNION ALL
SELECT 'Job Posting Tags', COUNT(*) FROM job_posting_tags
UNION ALL
SELECT 'Job Seeker Tags', COUNT(*) FROM job_seeker_tags
UNION ALL
SELECT 'Interships', COUNT(*) FROM interships
UNION ALL
SELECT 'Followup Milestones', COUNT(*) FROM followup_milestones
UNION ALL
SELECT 'Followup Issues', COUNT(*) FROM followup_issues
UNION ALL
SELECT 'Requests', COUNT(*) FROM requests
UNION ALL
SELECT 'Saved Jobs', COUNT(*) FROM saved_jobs
ORDER BY table_name;
