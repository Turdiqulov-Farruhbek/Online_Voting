CREATE TABLE election(
    id UUID PRIMARY KEY,
    name VARCHAR(256),
    open_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE candidate(
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES election(id),
    public_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);
ALTER TABLE 
    candidate 
ADD CONSTRAINT 
    unique_candidate_election_puclic UNIQUE(election_id, public_id);

CREATE TABLE public_vote(
    id UUID PRIMARY KEY,
    election_id UUID REFERENCES election(id) NOT NULL,
    public_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


ALTER TABLE 
    public_vote 
ADD CONSTRAINT 
    unique_public_vote_election_puclic UNIQUE(election_id, public_id);


CREATE TABLE vote (
    id UUID PRIMARY KEY,
    candidate_id UUID REFERENCES candidate(id) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);