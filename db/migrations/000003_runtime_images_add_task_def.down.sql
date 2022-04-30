-- Drop columns: taskdef, revision
-- Add columns: url, tag

ALTER TABLE runtime_images
    DROP COLUMN IF EXISTS `revision`;

ALTER TABLE runtime_images
    DROP COLUMN IF EXISTS `taskdef`;

ALTER TABLE runtime_images
    ADD COLUMN IF NOT EXISTS `language` VARCHAR(255)
        AFTER `language_name`;

ALTER TABLE runtime_images
    ADD COLUMN IF NOT EXISTS `tag` SMALLINT
        AFTER `language_name`;