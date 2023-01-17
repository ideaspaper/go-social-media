CREATE TABLE "users_tab" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" VARCHAR UNIQUE NOT NULL,
    "password" VARCHAR NOT NULL,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP
);
INSERT INTO "users_tab" (
        "email",
        "password",
        "first_name",
        "last_name",
        "created_at",
        "updated_at"
    )
VALUES (
        'acong@mail.com',
        '$2a$10$mTuOq/GlcQUPMmGGhogSR.Cgdh9D./6qRcSlK9.cRkSnoajjInTKq',
        'Acong',
        'Suherman',
        NOW(),
        NOW()
    ),
    (
        'djoko@mail.com',
        '$2a$10$mTuOq/GlcQUPMmGGhogSR.Cgdh9D./6qRcSlK9.cRkSnoajjInTKq',
        'Djoko',
        'Susanto',
        NOW(),
        NOW()
    );