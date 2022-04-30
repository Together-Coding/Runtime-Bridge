-- Drop columns: url, tag
-- Add columns: taskdef, revision

ALTER TABLE runtime_images
    DROP COLUMN IF EXISTS `url`;

ALTER TABLE runtime_images
    DROP COLUMN IF EXISTS `tag`;

ALTER TABLE runtime_images
    ADD COLUMN IF NOT EXISTS `taskdef` VARCHAR(255)
        AFTER language_name;

ALTER TABLE runtime_images
    ADD COLUMN IF NOT EXISTS `revision` SMALLINT
        AFTER taskdef;