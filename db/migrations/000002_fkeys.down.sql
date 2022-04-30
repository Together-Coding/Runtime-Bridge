ALTER TABLE runtime_allocations
    DROP FOREIGN KEY IF EXISTS `fk_rt_alloc_user`;

ALTER TABLE runtime_allocations
    DROP FOREIGN KEY IF EXISTS `fk_rt_alloc_rt_image`;

ALTER TABLE runtime_images
    DROP FOREIGN KEY IF EXISTS `fk_rt_image_lang`;

DROP INDEX IF EXISTS ix_rt_alloc_user_id ON runtime_allocations;
DROP INDEX IF EXISTS ix_rt_alloc_ct_api_key ON runtime_allocations;