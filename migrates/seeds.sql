-- SQL SEEDS
CREATE extension "uuid-ossp";
INSERT INTO public.events (id, name, properties, created_on)
    VALUES (
        (SELECT uuid_generate_v4()),
        'test_event',
        '{"foo":"bar"}',
        NOW()
    );