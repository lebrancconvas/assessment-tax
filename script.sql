CREATE TABLE IF NOT EXISTS "deductions" (
    "deduction_id" serial4 UNIQUE PRIMARY KEY NOT NULL,
    "personal_deduction" float,
    "kreceipt_deduction" float,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "used_flg" bool DEFAULT true
);
