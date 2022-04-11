ALTER TABLE runtime_allocations
    ADD CONSTRAINT `fk_rt_alloc_user`
        FOREIGN KEY IF NOT EXISTS (user_id) REFERENCES users (id)
            ON DELETE CASCADE;

ALTER TABLE runtime_allocations
    ADD CONSTRAINT `fk_rt_alloc_rt_image`
        FOREIGN KEY IF NOT EXISTS (user_id) REFERENCES users (id)
            ON DELETE CASCADE;

ALTER TABLE runtime_images
    ADD CONSTRAINT `fk_rt_image_lang`
        FOREIGN KEY IF NOT EXISTS (language_name) REFERENCES supported_languages (name)
            ON DELETE CASCADE;

CREATE INDEX
    IF NOT EXISTS ix_rt_alloc_user_id
    ON runtime_allocations (user_id);

CREATE INDEX
    IF NOT EXISTS ix_rt_alloc_ct_api_key
    ON runtime_allocations (cont_api_key);

