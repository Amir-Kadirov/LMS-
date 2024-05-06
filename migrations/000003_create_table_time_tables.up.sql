CREATE TABLE IF NOT EXISTS  time_tables (
    id uuid PRIMARY KEY,
    teacher_id uuid NOT NULL,
    student_id uuid NOT NULL,
    subject_id uuid NOT NULL,
    start_date timestamp NOT NULL,
    end_date timestamp NOT NULL
);

ALTER TABLE time_tables ADD FOREIGN KEY (teacher_id) REFERENCES teacher (id);
ALTER TABLE time_tables ADD FOREIGN KEY (student_id) REFERENCES students (id);
ALTER TABLE time_tables ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);