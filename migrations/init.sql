CREATE TYPE device_state AS ENUM ('available', 'in-use', 'inactive');

CREATE TABLE IF NOT EXISTS devices (
    id              BIGSERIAL       PRIMARY KEY,
    external_id     UUID            NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    name            VARCHAR(255)    NOT NULL,
    brand           VARCHAR(255)    NOT NULL,
    state           device_state    NOT NULL DEFAULT 'available',
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ,
    is_deleted      BOOLEAN         NOT NULL DEFAULT FALSE,
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_devices_brand ON devices(brand) WHERE is_deleted = FALSE;
CREATE INDEX IF NOT EXISTS idx_devices_state ON devices(state) WHERE is_deleted = FALSE;
-- Test devices
INSERT INTO devices (external_id, name, brand, state) 
VALUES ('77b85ecb-e767-4f2d-a3cf-be53dae49274','Test device 1', 'Brand number one', 'available');
INSERT INTO devices (external_id, name, brand, state) 
VALUES ('f3b4ea8f-1e69-4736-9cc7-f8775761fefd','Test device 2', 'Brand number one', 'in-use');
INSERT INTO devices (external_id, name, brand, state) 
VALUES ('c343dabd-0afd-4f5b-a632-42be259df112','Test device 3', 'Brand number two', 'inactive');
INSERT INTO devices (external_id, name, brand, state) 
VALUES ('881ff09c-4ac4-43e7-9301-263e0f930273','Test device 4', 'Brand number two', 'available');
INSERT INTO devices (external_id, name, brand, state) 
VALUES ('4e579ef0-0cd8-459c-bfa5-3470802bab0c','Test device 5', 'Brand number three', 'available');


/* 
    Triggers to enforce business rules
*/
CREATE OR REPLACE FUNCTION prevent_created_at_update()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.created_at IS DISTINCT FROM OLD.created_at THEN
        RAISE EXCEPTION 'created_at cannot be updated after creation';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER trg_prevent_created_at_update BEFORE UPDATE ON devices FOR EACH ROW EXECUTE FUNCTION prevent_created_at_update();

CREATE OR REPLACE FUNCTION prevent_device_inuse_update()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.state = 'in-use' THEN
        RAISE EXCEPTION 'Cannot update device while it is in use';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER trg_prevent_device_inuse_update BEFORE UPDATE ON devices FOR EACH ROW EXECUTE FUNCTION prevent_device_inuse_update();

CREATE OR REPLACE FUNCTION handle_device_deletion()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.state = 'in-use' AND NEW.is_deleted = TRUE THEN
        RAISE EXCEPTION 'Cannot delete device while it is in use';
    END IF;

    IF NEW.is_deleted = TRUE AND OLD.is_deleted = FALSE THEN
        NEW.deleted_at := CURRENT_TIMESTAMP;
    ELSIF NEW.is_deleted = FALSE AND OLD.is_deleted = TRUE THEN
        NEW.deleted_at := NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_handle_device_deletion BEFORE UPDATE ON devices FOR EACH ROW EXECUTE FUNCTION handle_device_deletion();
